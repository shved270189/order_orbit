import Button from 'react-bootstrap/Button';
import Col from 'react-bootstrap/Col';
import Form from 'react-bootstrap/Form';
import Row from 'react-bootstrap/Row';
import axios from 'axios';
import { useNavigate } from "react-router";
import { useState } from 'react'
import BackButton from '../common/BackButton';

type User = {
  fullName: string;
  login: string;
};

function User() {
  const [userAttrs, setUserAttrs] = useState<User>({ fullName: '', login: '' });
  const navigate = useNavigate();
  const createUser = async () => {
    console.log('User Attrs:', userAttrs);
    try {
      const response = await axios.post(`${import.meta.env.VITE_API_BASE_URL}/users`, userAttrs);
      console.log('User created:', response.data);
      navigate(`/users/${response.data.id}`);
    } catch (error) {
      console.error('Error creating user:', error);
    }
  };

  return (
    <>
    <BackButton to="/users" />
    <Form>
      <Row className="mb-3">
        <Form.Group as={Col} controlId="formGridLogin">
          <Form.Label>Login</Form.Label>
          <Form.Control type="text" placeholder="Enter login" onChange={(e) => setUserAttrs({ ...userAttrs, login: e.target.value })} />
        </Form.Group>

        <Form.Group as={Col} controlId="formGridFullName">
          <Form.Label>Full Name</Form.Label>
          <Form.Control type="text" placeholder="Enter full name" onChange={(e) => setUserAttrs({ ...userAttrs, fullName: e.target.value })} />
        </Form.Group>
      </Row>
      <Button variant="primary" onClick={createUser}>
        Create
      </Button>
    </Form>
    </>
  )
}

export default User
