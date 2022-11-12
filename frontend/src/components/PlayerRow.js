import SeasonRow from "./SeasonRow";
import DropdownButton from "./DropdownButton";
import React, { useState } from 'react';
import { seasonToFilter } from '../util/filters'
import { getImage } from '../util/util'

function PlayerRow(props) {
    const [dropdown, setDropdown] = useState(false);
    const fullName = props.firstName + " " + props.lastName;

    if (props.cardData.length > 0) {
        return (
            <>
                <tr>
                    <td><img src={getImage(props.firstName, props.lastName)} alt={fullName}></img></td>
                    <td>{props.firstName}</td>
                    <td>{props.lastName}</td>
                    <td>{props.seasonsPlayed}</td>
                    <td>{props.cardData.length}</td>
                    <td><DropdownButton state={dropdown} setState={setDropdown} /></td>
                </tr>
                {dropdown &&
                    <>
                        {props.seasons.map(season => (
                            <SeasonRow season={season} cardData={props.cardData.filter(seasonToFilter(season))} />
                        ))
                        }
                    </>
                }
            </>
        );
    }
}

export default PlayerRow;