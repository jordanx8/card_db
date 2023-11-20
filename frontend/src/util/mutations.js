import { gql } from "@apollo/client";

export const ADD_PLAYER_MUTATION = gql`
  mutation AddPlayer($firstName: String, $lastName: String, $seasonsPlayed: String) {
    AddPlayer(firstName: $firstName, lastName: $lastName, seasonsPlayed: $seasonsPlayed)
  }
`;

export const ADD_SEASON_MUTATION = gql`
  mutation AddSeason($firstName: String, $lastName: String, $season: String) {
    AddSeason(firstName: $firstName, lastName: $lastName, season: $season)
  }
`;

export const ADD_CARD_MUTATION = gql`
  mutation AddCard($playerName: String, $season: String, $manufacturer: String, $set: String, $insert: String, $parallel: String, $cardNumber: String, $notes: [String], $imageLink: String, $tableName: String) {
    AddCard(playerName: $playerName, season: $season, manufacturer: $manufacturer, set: $set, insert: $insert, parallel: $parallel, cardNumber: $cardNumber, notes: $notes, imageLink: $imageLink, tableName: $tableName)
  }
`;