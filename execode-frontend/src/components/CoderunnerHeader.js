import PropTypes from 'prop-types'
import CoderunnerInput from './CoderunnerInput';

const CoderunnerHeader = (props) => {
    const codeRunnerHeadingStyle = {
        display: "flex",
        padding: "20px 20px",
        alignItems: "center",
        justifyContent: "flex-start"
    };

    return (
    <header style={codeRunnerHeadingStyle}>
        <CoderunnerInput />
        <div>
            <label>Language: </label>
            <select name="languages" id="languages">
                <option value="bash">Bash</option>
                <option value="c">C</option>
                <option value="c++">C++</option>
                <option value="csharp">C#</option>
                <option value="dart">Dart</option>
                <option value="go">Go</option>
                <option value="java">Java</option>
                <option value="javascript" selected>JavaScript</option>
                <option value="kotlin">Kotlin</option>
                <option value="lua">Lua</option>
                <option value="php">php</option>
                <option value="python">Python 3</option>
                <option value="python2">Python 2</option>
                <option value="ruby">Ruby</option>
                <option value="rust">Rust</option>
                <option value="scala">Scala</option>
                <option value="sqlite3">SQLite3</option>
                <option value="swift">Swift</option>
                <option value="typescript">TypeScript</option>
            </select>
        </div>
    </header>
  )
}

CoderunnerHeader.propTypes = {
    title: PropTypes.string
}

export default CoderunnerHeader