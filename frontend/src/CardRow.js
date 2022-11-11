function CardRow(props) {
    return (
        <tr>
            <td>{props.season}</td>
            <td>{props.manufacturer}</td>
            <td>{props.set}</td>
            <td>{props.insert}</td>
            <td>{props.parallel}</td>
            <td>{props.cardNumber}</td>
            <td>{props.notes.map(note =>
                (<div>{note}</div>))}</td>
        </tr>
    );
}

export default CardRow;
