import React, { useState } from 'react';
import PlayerRow from "./PlayerRow";
import { v4 as uuidv4 } from 'uuid';
import Table from 'react-bootstrap/Table';
import InputGroupHandler from './InputGroupHandler';

function CardTable(props) {
  const [cards, setCards] = useState(props.cardData.default);
  const [players, setPlayers] = useState(props.playerData.default);
  const [checked, setChecked] = useState(false);
  const [originalData, setOriginalData] = useState([]) 

  function nameToFilter(name) {
    return function filterCardsByName(card) {
      return card.playerName === name;
    }
  }

  return (
    <React.Fragment key={uuidv4()}>
      <h1>{props.tableName}</h1>
      <InputGroupHandler originalData={originalData} setOriginalData={setOriginalData} checked={checked} setChecked={setChecked} cards={cards} setCards={setCards} players={players} setPlayers={setPlayers} />
      <Table striped bordered hover responsive variant="dark">
        <thead>
          <tr>
            <th>Picture:</th>
            <th>First Name:</th>
            <th>Last Name:</th>
            <th>Seasons Played:</th>
            <th>Number of Cards:</th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          {players.map(values => (
            <PlayerRow firstName={values.firstName} lastName={values.lastName} seasonsPlayed={values.seasonsPlayed} seasons={values.seasons} cardData={cards.filter(nameToFilter((values.firstName + " " + values.lastName)))} />
          ))
          }
        </tbody>
      </Table>
    </React.Fragment>
  );
}

export default CardTable;
