import React from 'react';
import ReactDOM from 'react-dom/client';
import CardTable from './CardTable';

import * as playerData from './test_player_data.json';
import * as cardData from './test_card_data.json';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <CardTable cardData={cardData} playerData={playerData}/>
  </React.StrictMode>
);
