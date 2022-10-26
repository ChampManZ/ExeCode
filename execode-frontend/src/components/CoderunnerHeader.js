import PropTypes from 'prop-types'
import CoderunnerInput from './CoderunnerInput';
import Select from 'react-select';

const CoderunnerHeader = (props) => {

    const codeRunnerHeadingStyle = {
        display: "flex",
        padding: "20px 20px",
        alignItems: "center",
        justifyContent: "flex-start"
    };

const langOptions = [
    { value: 'bash', label: 'Bash' },
    { value: 'c', label: 'C' },
    { value: 'c++', label: 'C++' },
    { value: 'csharp', label: 'C#' },
    { value: 'dart', label: 'Dart' },
    { value: 'go', label: 'Go' },
    { value: 'java', label: 'Java' },
    { value: 'javascript', label: 'JavaScript' },
    { value: 'kotlin', label: 'Kotlin' },
    { value: 'lua', label: 'Lua' },
    { value: 'php', label: 'php' },
    { value: 'python', label: 'Python3' },
    { value: 'python2', label: 'Python2' },
    { value: 'ruby', label: 'Ruby' },
    { value: 'rust', label: 'Rust' },
    { value: 'scala', label: 'Scala' },
    { value: 'sqlite3', label: 'SQLite3' },
    { value: 'swift', label: 'Swift' },
    { value: 'typescript', label: 'TypeScript' },
]

    return (
    <header style={codeRunnerHeadingStyle}>
        <CoderunnerInput />
        <div id='lang-dropdown'>
            <label>Language: </label>
            <Select value={langOptions.value} options={langOptions} defaultValue={langOptions[7]} />
        </div>
        <hr />
    </header>
  )
}

CoderunnerHeader.propTypes = {
    title: PropTypes.string
}

export default CoderunnerHeader