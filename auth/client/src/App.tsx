import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

function App() {

    const handleLoginDiscord = () => {
        window.location.href = "http://localhost:3000/auth/discord";
    };
    const handleLoginGoogle = () => {
        window.location.href = "http://localhost:3000/auth/google";
    };

    return (
        <>
            <div>
                <a href="https://vitejs.dev" target="_blank">
                    <img src={viteLogo} className="logo" alt="Vite logo"/>
                </a>
                <a href="https://react.dev" target="_blank">
                    <img src={reactLogo} className="logo react" alt="React logo"/>
                </a>
            </div>
            <h1>Vite + React</h1>
            <div>
                <button onClick={handleLoginGoogle}>
                    Login with Google
                </button>
            </div>

            <div>
                <button onClick={handleLoginDiscord}>
                    Login with Discord
                </button>
            </div>
            <p className="read-the-docs">
                Click on the Vite and React logos to learn more
            </p>
        </>
    )
}

export default App