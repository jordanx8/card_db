import DropdownButton from "./DropdownButton";
import CardRow from "./CardRow";
import React, { useState } from 'react';
import { v4 as uuidv4 } from 'uuid';

function SeasonRow(props) {
    const [dropdown, setDropdown] = useState(false);
    return (
        <React.Fragment key={uuidv4()}>
            <tr>
                <td>{props.season}</td>
                <td><DropdownButton state={dropdown} setState={setDropdown} /></td>
            </tr>
            {dropdown &&
                <tr>
                    <td>
                        <table>
                            <tbody>
                                <tr>
                                    <td>Season</td>
                                    <td>Manufacturer</td>
                                    <td>Set</td>
                                    <td>Insert</td>
                                    <td>Parallel</td>
                                    <td>Card Number</td>
                                    <td>Notes</td>
                                </tr>
                                {props.cardData.map(card => (
                                    <CardRow name={card.playerName}
                                        season={card.season}
                                        manufacturer={card.manufacturer}
                                        set={card.set}
                                        insert={card.insert}
                                        parallel={card.parallel}
                                        cardNumber={card.cardNumber}
                                        notes={card.notes}
                                    />
                                ))
                                }
                            </tbody>
                        </table>
                    </td>
                </tr>
            }
        </React.Fragment>
    );
}

export default SeasonRow;
