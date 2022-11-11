import { gql } from "@apollo/client";

export const GET_CARDS = gql`
  query cards {
    cards {
        FirstName
        LastName
        SeasonsPlayed
        Season
        Manufacturer
        Set
        Insert
        Parallel
        CardNumber
        Notes
    }
  }
`;