import { StrictMode } from 'react'
import ReactDOM from "react-dom/client";
import { BrowserRouter, Routes, Route, NavLink } from "react-router";
import 'bootstrap/dist/css/bootstrap.min.css';
import Home from './components/Home.tsx'
import Users from './components/Users.tsx'
import Container from 'react-bootstrap/esm/Container';

ReactDOM.createRoot(document.getElementById('root')!).render(
  <StrictMode>
      <BrowserRouter>
        <nav className="navbar navbar-expand-lg bg-body-tertiary">
          <div className="container-fluid">
            <NavLink className="navbar-brand" to="/" end>
              Order Orbit
            </NavLink>
            <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
              <span className="navbar-toggler-icon"></span>
            </button>
            <div className="collapse navbar-collapse" id="navbarSupportedContent">
              <ul className="navbar-nav me-auto mb-2 mb-lg-0">
                <li className="nav-item">
                  <NavLink className="nav-link active" to="/users" end>
                    Users
                  </NavLink>
                </li>
              </ul>
            </div>
          </div>
        </nav>
        <Container>
          <Routes>
            <Route index element={<Home />} />

            <Route path="users">
              <Route index element={<Users />} />
              {/* <Route path="new" element={<City />} />
              <Route path=":id" element={<City />} />
              <Route path=":id/edit" element={<City />} /> */}
            </Route>
          </Routes>
        </Container>
      </BrowserRouter>
  </StrictMode>,
)
