import { gql } from "@apollo/client";

export const ADD_PLAYER_MUTATION = gql`
  mutation AddPlayer($firstName: String, $lastName: String, $seasonsPlayed: String) {
    AddPlayer(firstName: $firstName, lastName: $lastName, seasonsPlayed: $seasonsPlayed)
  }
`;