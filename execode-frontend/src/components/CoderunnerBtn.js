function CoderunnerBtn(props) {
    const btnStyle = {
        textAlign: "right"
      }

    return (
        <div style={btnStyle}>
            <button type="button">{props.name}</button>
        </div>
    )
}

export default CoderunnerBtn