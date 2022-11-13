import Image from 'react-bootstrap/Image'

function CardImageColumn(props) {
    return (
        <td><Image fluid width="99" height="140" src={props.image} /></td>
    )
}

export default CardImageColumn