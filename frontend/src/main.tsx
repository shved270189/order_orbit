import { StrictMode } from 'react'
import ReactDOM from "react-dom/client";
import { BrowserRouter, Routes, Route } from "react-router";
import 'bootstrap/dist/css/bootstrap.min.css'
import App from './App.tsx'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <BrowserRouter>
      <Routes>
        <Route index element={<App />} />

        <Route path="users">
          <Route index element={<App />} />
          {/* <Route path="new" element={<City />} />
          <Route path=":id" element={<City />} />
          <Route path=":id/edit" element={<City />} /> */}
        </Route>
      </Routes>
    </BrowserRouter>
  </StrictMode>,
)
