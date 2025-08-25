import { StrictMode } from 'react'
import ReactDOM from "react-dom/client";
import { BrowserRouter} from "react-router";

import Header from './components/layout/Header.tsx'
import Content from './components/layout/Content.tsx'

import 'bootstrap/dist/css/bootstrap.min.css';


ReactDOM.createRoot(document.getElementById('root')!).render(
  <StrictMode>
      <BrowserRouter>
        <Header />
        <Content />
      </BrowserRouter>
  </StrictMode>,
)
