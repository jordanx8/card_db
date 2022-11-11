function DropdownButton(props) {
    function useStateHandler(){
        props.setState(!props.state)
    }

    return (
        <button onClick={useStateHandler}>^</button>
    );
}

export default DropdownButton;
