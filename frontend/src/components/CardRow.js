import CardImageColumn from "./CardImageColumn";

function CardRow({ manufacturer, set, insert, parallel, cardNumber, notes, image }) {
    return (
        <>
            <tr>
                <CardImageColumn image={image} />
                <td>{manufacturer}</td>
                <td>{set}</td>
                <td>{insert}</td>
                <td>{parallel}</td>
                <td>{cardNumber}</td>
                <td>{notes.map(note =>
                    (<div>{note}</div>))}</td>
            </tr>
        </>
    );
}

export default CardRow;
