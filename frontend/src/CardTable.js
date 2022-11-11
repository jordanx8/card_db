import React from "react";
import PlayerRow from "./PlayerRow";
import { v4 as uuidv4 } from 'uuid';

function CardTable(props) {
  function nameToFilter(name) {
    return function filterCardsByName(card) {
      return card.playerName === name;
    }
  }


  return (
    <React.Fragment key={uuidv4()}>
      <h1>Card Table:</h1>
      <table>
        <tbody>
          <tr>
            <th>Picture:</th>
            <th>First Name:</th>
            <th>Last Name:</th>
            <th>Seasons Played:</th>
          </tr>
          {props.playerData.default.map(values => (
              <PlayerRow firstName={values.firstName} lastName={values.lastName} seasonsPlayed={values.seasonsPlayed} seasons={values.seasons} cardData={props.cardData.default.filter(nameToFilter((values.firstName + " " + values.lastName)))} />
          ))
          }

        </tbody>
      </table>
    </React.Fragment>
  );
}

export default CardTable;
