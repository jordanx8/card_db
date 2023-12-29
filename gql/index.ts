const { ApolloServer, gql } = require('apollo-server');
const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const PROTO_PATH = 'cards.proto';
const packageDefinition = protoLoader.loadSync(
  PROTO_PATH,
  {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
    includeDirs: ['protos']
  });
const proto = grpc.loadPackageDefinition(packageDefinition).main;
const client = new proto.CardService('localhost:8080', grpc.credentials.createInsecure());

const typeDefs = gql`
  type Player {
    firstName: String
    lastName: String
    seasonsPlayed: String
    seasons: [String]
  }

  type Card {
    playerName: String
    season: String
    manufacturer: String
    set: String
    insert: String
    parallel: String
    cardNumber: String
    notes: [String]
    imageLink: String
    team: String
  }

  type Query {
    cards(tableName: String, search: String): [Card]
    players(empty: Int): [Player]
  }

  type Mutation {
    AddPlayer(firstName: String, lastName: String, seasonsPlayed: String): String
    AddSeason(firstName: String, lastName: String, season: String): String
    AddCard(playerName: String, season: String, manufacturer: String, set: String, insert: String, parallel: String, cardNumber: String, notes: [String], imageLink: String, tableName: String, team: String): String
  }
`;

const resolvers = {
  Query: {
    cards: (parent, args, context, info) => {
      return new Promise((resolve) => {
        let call = client.GetCards({
          tableName: args.tableName,
          search: args.search
        });
        let cards = [];
        call.on('data', function(card) {
          cards.push(card);
        });
        call.on('end', function() {
          resolve(cards);
        });
      }
    )},
    players: () => {
      return new Promise((resolve) => {
        let call = client.GetPlayers({});
        let players = [];
        call.on('data', function(player) {
          players.push(player);
        });
        call.on('end', function() {
          resolve(players);
        });
      }
    )},
  },
    Mutation: {
      AddPlayer: (parent, args, context, info) => {
        return new Promise((resolve) => client.AddPlayer({
          firstName: args.firstName, 
          lastName: args.lastName,
          seasonsPlayed: args.seasonsPlayed}, function (err, response) {
          if (err) {
            resolve(err.details)
          } else {
            resolve(response.response)
          }
        }));
      },
      AddSeason: (parent, args, context, info) => {
        return new Promise((resolve) => client.AddSeason({
          firstName: args.firstName, 
          lastName: args.lastName,
          season: args.season}, function (err, response) {
          if (err) {
            resolve(err.details)
          } else {
            resolve(response.response)
          }
        }));
      },
      AddCard: (parent, args, context, info) => {
        console.log(args)
        return new Promise((resolve) => client.AddCard({
          playerName: args.playerName, 
          season: args.season,
          manufacturer: args.manufacturer,         
          set: args.set, 
          insert: args.insert,
          parallel: args.parallel,          
          cardNumber: args.cardNumber, 
          notes: args.notes,
          imageLink: args.imageLink,
          tableName: args.tableName,
          team: args.team}, function (err, response) {
          if (err) {
            resolve(err.details)
          } else {
            resolve(response.response)
          }
        }));
      },
    },
  }

const server = new ApolloServer({
  typeDefs,
  resolvers
});


server.listen('4000').then(({ url }) => {
  console.log(`🚀 Server ready at ${url}`)
});