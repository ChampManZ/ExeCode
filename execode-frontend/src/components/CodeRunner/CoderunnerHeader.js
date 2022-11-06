import CoderunnerInput from './CoderunnerInput';
import Select from 'react-select';

const CoderunnerHeader = () => {

    const codeRunnerHeadingStyle = {
        display: "flex",
        padding: "20px 20px",
        alignItems: "center",
        justifyContent: "flex-start",
        borderBottom: "1px solid #E5E5E5"
    };

    // langOptions will be use to display dropdown available programming language
    // Addition, we will use React Select library to enhance our code/workflow
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
    </header>
  )
}

export default CoderunnerHeader