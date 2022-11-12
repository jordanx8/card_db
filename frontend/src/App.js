import React from 'react';
import * as playerData from './test_player_data.json';
import * as cardData from './test_card_data.json';
import CardTable from './components/CardTable';
import Navigation from './components/Navigation';

function App() {
    return (
        <>
            <Navigation />
            <CardTable tableName={"New Orleans Pelicans"} cardData={cardData} playerData={playerData} />
        </>
    );
}

export default App;