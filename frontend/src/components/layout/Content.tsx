import { Routes, Route } from "react-router";
import Container from 'react-bootstrap/esm/Container';

import Home from './../Home.tsx'
import Users from '../users/Users.tsx'
import User from '../users/User.tsx'
import CreateUserForm from '../users/CreateUserForm.tsx'

import './Content.css'

function Content() {
  return (
    <Container id="content-container">
      <Routes>
        <Route index element={<Home />} />
        <Route path="users/new" element={<CreateUserForm />} />
        <Route path="users/:userId" element={<User />} />
        <Route path="users" element={<Users />} />
      </Routes>
    </Container>
  )
}
export default Content
