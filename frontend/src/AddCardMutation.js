import { gql } from "@apollo/client";

export const ADD_CARD = gql`
  mutation Addcard($firstName: String, $lastName: String, $seasonsPlayed: String, $season: String, $manufacturer: String, $set: String, $insert: String, $parallel: String, $cardNumber: String, $notes: String) {
    addcard(FirstName: $firstName, LastName: $lastName, SeasonsPlayed: $seasonsPlayed, Season: $season, Manufacturer: $manufacturer, Set: $set, Insert: $insert, Parallel: $parallel, CardNumber: $cardNumber, Notes: $notes)
  }
`;