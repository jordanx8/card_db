import DropdownButton from "./DropdownButton";
import CardRow from "./CardRow";
import React, { useState } from 'react';

function SeasonRow({ season, cardData }) {
    const [dropdown, setDropdown] = useState(false);
    if (cardData.length > 0) {
        return (
            <>
                <tr>
                    <th colSpan={6}>{season}</th>
                    <th><DropdownButton state={dropdown} setState={setDropdown} /></th>
                </tr>
                {dropdown &&
                    <>
                        <tr>
                            <th>Image</th>
                            <th>Manufacturer</th>
                            <th>Set</th>
                            <th>Insert</th>
                            <th>Parallel</th>
                            <th>Card Number</th>
                            <th>Notes</th>
                        </tr>
                        {cardData.map(card => (
                            <CardRow
                                manufacturer={card.manufacturer}
                                set={card.set}
                                insert={card.insert}
                                parallel={card.parallel}
                                cardNumber={card.cardNumber}
                                notes={card.notes}
                                image={card.imageLink}
                            />
                        ))
                        }
                    </>
                }
            </>
        );
    }
}

export default SeasonRow;
