import {useState} from 'react'
import './App.css'

function App() {
    const [count, setCount] = useState(0)

    return (
        <div>
            <h1>"Привет!!!"</h1>
            <button onClick={async () => {
                const response = await fetch("http://localhost:8000/app")
                const data = await response.json()
                console.log(data)
            }}>Click
            </button>
        </div>
    )
}

export default App
