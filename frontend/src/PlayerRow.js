import SeasonRow from "./SeasonRow";
import DropdownButton from "./DropdownButton";
import React, { useState } from 'react';
import { v4 as uuidv4 } from 'uuid';

function PlayerRow(props) {
    const [dropdown, setDropdown] = useState(false);
    const fullName = props.firstName + " " + props.lastName;
    function getImage(firstName, lastName) {
        let num = "01"
        const twonames = ["Trey_Murphy", "Larry_Nance", "Cameron_Thomas", "Anthony_Davis"]
        for (let i = 0; i < twonames.length; i++) {
            if (firstName + "_" + lastName.replace("-", "_") === twonames[i]) {
                num = "02"
            }
        }
        return "https://www.basketball-reference.com/req/202106291/images/players/" + lastName.toLowerCase().slice(0, 5) + firstName.toLowerCase().slice(0, 2) + num + ".jpg"
    }

    function seasonToFilter(season) {
        return function filterCardsBySeason(card) {
            return season.includes(card.season);
        }
    }

    return (
        <React.Fragment key={uuidv4()}>
            <tr>
                <td><img src={getImage(props.firstName, props.lastName)} alt={fullName}></img></td>
                <td>{props.firstName}</td>
                <td>{props.lastName}</td>
                <td>{props.seasonsPlayed}</td>
                <td>{props.cardData.length}</td>
                <td><DropdownButton state={dropdown} setState={setDropdown} /></td>
            </tr>
            {dropdown &&
                <React.Fragment key={uuidv4()}>
                    {props.seasons.map(season => (
                        <SeasonRow season={season} cardData={props.cardData.filter(seasonToFilter(season))} />
                    ))
                    }
                </React.Fragment>
            }
        </React.Fragment>
    );
}

export default PlayerRow;