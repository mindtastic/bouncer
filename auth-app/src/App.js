import logo from './logo.svg';
import './App.css';
import { ReactKeycloakProvider } from "@react-keycloak/web";
import keycloak from "./Keycloak"
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Nav from "./components/Nav";
import Homepage from "./pages/Homepage";
import PrivateRoute from "./helpers/PrivateRoute";


function App() {
  return (
    <div className="App">
      <ReactKeycloakProvider authClient={keycloak}>
        <Nav />
        {/* <BrowserRouter>
          <Routes>
            <Route
              path="/secured"
              element={
                <PrivateRoute>
                  <Homepage />
                </PrivateRoute>
              }
            />
          </Routes>
        </BrowserRouter> */}

        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />

          <a
            className="App-link"
            href="https://reactjs.org"
            target="_blank"
            rel="noopener noreferrer"
          >
            Hello Mindtatsic
          </a>
        </header>

      </ReactKeycloakProvider>

    </div>
  );
}

export default App;
