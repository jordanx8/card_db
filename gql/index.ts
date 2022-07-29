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
    firstname: String
    lastname: String
    seasonsplayed: String
  }

  type Card {
    firstname: String
    lastname: String
    seasonsplayed: String
    season: String
    manufacturer: String
    set: String
    insert: String
    parallel: String
    cardnumber: String
    notes: String
  }

  type Query {
    cards(empty: Int): [Card]
    players(empty: Int): [Player]
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
    players: (parent, args, context, info) => {
      return new Promise((resolve) => client.GetPlayers({}, function (err, response) {
        if (err == undefined) {
          resolve(response.players);
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