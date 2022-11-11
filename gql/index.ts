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
const client = new proto.CardService('127.0.0.1:8080', grpc.credentials.createInsecure());

const typeDefs = gql`
  type Player {
    firstname: String
    lastname: String
    seasonsplayed: String
  }

  type Card {
    FirstName: String
    LastName: String
    SeasonsPlayed: String
    Season: String
    Manufacturer: String
    Set: String
    Insert: String
    Parallel: String
    CardNumber: String
    Notes: String
  }

  type Query {
    cards(empty: Int): [Card]
    players(empty: Int): [Player]
  }

  type Mutation {
    addcard(FirstName: String, LastName: String, SeasonsPlayed: String, Season: String, Manufacturer: String, Set: String, Insert: String, Parallel: String, CardNumber: String, Notes: String): Int
  }
`;

const resolvers = {
  Query: {
    cards: (parent, args, context, info) => {
      console.log("cards")
      return new Promise((resolve) => client.GetCards({}, function (err, response) {
        if (err == undefined) {
          resolve(response.cards);
        }
      }));
    },
    players: (parent, args, context, info) => {
      return new Promise((resolve) => client.GetPlayers({}, function (err, response) {
        if (err == undefined) {
          resolve(response.players);
        }
      }));
    },
  },
    Mutation: {
      addcard: (parent, args, context, info) => {
        console.log(args)
        return new Promise((resolve) => client.AddCard({FirstName: args["FirstName"], LastName: args["LastName"], SeasonsPlayed: args["SeasonsPlayed"], Season: args["Season"], Manufacturer: args["Manufacturer"], Set: args["Set"], Insert: args["Insert"], Parallel: args["Parallel"], CardNumber: args["CardNumber"], Notes: args["Notes"]}, function (err, response) {
          if (err == undefined) {
            resolve(response.success);
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