function CoderunnerBtn(props) {
    return (
        <div>
            <button type="button" style={props.style} onClick={props.clickFunc} >{props.name}</button>
        </div>
    )
}

export default CoderunnerBtn