import Button from 'react-bootstrap/Button';

function DropdownButton(props) {
    function useStateHandler(){
        props.setState(!props.state)
    }

    if(!props.state)
    {
        return (
            <Button onClick={useStateHandler}>Expand</Button>
        );
    }else{
        return (
            <Button onClick={useStateHandler}>Collapse</Button>
        );
    }

}

export default DropdownButton;
