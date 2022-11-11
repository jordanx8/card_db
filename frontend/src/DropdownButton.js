import { ArrowUp, ArrowDown } from 'react-bootstrap-icons';

function DropdownButton(props) {
    function useStateHandler(){
        props.setState(!props.state)
    }

    if(!props.state)
    {
        return (
            <button onClick={useStateHandler}><ArrowUp/></button>
        );
    }else{
        return (
            <button onClick={useStateHandler}><ArrowDown/></button>
        );
    }

}

export default DropdownButton;
