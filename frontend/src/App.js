import React from 'react';
import CardTable from './components/CardTable';
import Navigation from './components/Navigation';

function App() {
    return (
        <>
            <Navigation />
            <CardTable tableName={"New Orleans Pelicans"} />
        </>
    );
}

export default App;