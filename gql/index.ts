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
  }

  type Query {
    cards(empty: Int): [Card]
    players(empty: Int): [Player]
  }

  type Mutation {
    AddPlayer(firstName: String, lastName: String, seasonsPlayed: String, seasons: [String]): Boolean
  }
`;

const resolvers = {
  Query: {
    cards: (parent, args, context, info) => {
      return new Promise((resolve) => client.GetCards({}, function (err, response) {
        if (err == undefined) {
          resolve(response.cards);
        }
      }));
    },
    players: () => {
      return new Promise((resolve) => {
        let call = client.GetPlayers({});
        let players = [];
        call.on('data', function(player) {
          players.push(player);
          return player
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
          seasonsPlayed: args.seasonsPlayed,
          seasons: args.seasons}, function (err, response) {
          if (err) {
            console.log(err)
            resolve(false)
          } else {
            resolve(response.success)
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