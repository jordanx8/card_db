import Image from 'react-bootstrap/Image'

function CardImageColumn({ image }) {
    return (
        <td><Image fluid width="99" height="140" src={image} /></td>
    )
}

export default CardImageColumn