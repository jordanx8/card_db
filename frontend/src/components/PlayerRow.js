import SeasonRow from "./SeasonRow";
import DropdownButton from "./DropdownButton";
import React, { useState } from 'react';
import { getImage, seasonToFilter } from '../util/util'

function PlayerRow({ firstName, lastName, seasons, seasonsPlayed, cardData }) {
    const [dropdown, setDropdown] = useState(false);
    const fullName = firstName + " " + lastName;

    if (cardData.length > 0) {
        return (
            <>
                <tr>
                    <td><img src={getImage(firstName, lastName)} alt={fullName}></img></td>
                    <td>{firstName}</td>
                    <td>{lastName}</td>
                    <td>{seasonsPlayed}</td>
                    <td colSpan={2}>{cardData.length}</td>
                    <td><DropdownButton state={dropdown} setState={setDropdown} /></td>
                </tr>
                {dropdown &&
                    <>
                        {seasons.map(season => (
                            <SeasonRow season={season} cardData={cardData.filter(seasonToFilter(season))} />
                        ))
                        }
                    </>
                }
            </>
        );
    }
}

export default PlayerRow;