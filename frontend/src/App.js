import React, { useState } from 'react';
import * as playerData from './test_player_data.json';
import * as cardData from './test_card_data.json';
import CardTable from './CardTable';
import Navigation from './Navigation';

function App() {
    const [searchString, setSearchString] = useState('');
    return (
        <>
            <Navigation />
            <CardTable tableName={"New Orleans Pelicans"} cardData={cardData} playerData={playerData} setSearchString={setSearchString} />
        </>
    );
}

export default App;