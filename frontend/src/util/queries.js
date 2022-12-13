import { gql } from "@apollo/client";

export const GET_PLAYERS_QUERY = gql`
  {
    players {
        firstName
        lastName
        seasons
        seasonsPlayed
    }
  }
`;

export const GET_CARDS_QUERY = gql`
  query Cards($tableName: String) {
    cards(tableName: $tableName) {
      cardNumber
      imageLink
      insert
      manufacturer
      notes
      parallel
      playerName
      season
      set
    }
  }
`;