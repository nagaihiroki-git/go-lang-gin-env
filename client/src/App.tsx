import React from 'react';
import logo from './logo.svg';
import './App.css';
import axios from "axios";


function App() {
  const [txt, setTxt] = React.useState<string>("");
  const [user, setUser] = React.useState<User>({name: ""});

  interface User {
   name: string;
  }
  const test = () => {
    console.log("test");
    const baseURL = "http://localhost:3000/test2?id="+txt;

    axios.get(baseURL).then((res) => {
      setUser(res.data);
    });
  }
  // React.useEffect(() => {
  //   axios.get<User>(baseURL).then((response) => {
  //     console.log(response.data);
  //     setUser(response.data);
  //   });
  // }, []);

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <input
            value={txt}
            type={"text"}
            onChange={event => setTxt(event.target.value)}/>
        <button onClick={()=>test()}>Click me</button>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          {user.name}
        </a>
      </header>
    </div>
  );
}

export default App;
