import Button from 'react-bootstrap/Button';
import Col from 'react-bootstrap/Col';
import Form from 'react-bootstrap/Form';
import Row from 'react-bootstrap/Row';
import axios from 'axios';
import { useNavigate } from "react-router";
import { useState, useEffect } from 'react'
import BackButton from '../common/BackButton';
import {useParams} from "react-router";

type User = {
  id: number;
  fullName: string;
  login: string;
  updatedAt: string;
  createdAt: string;
};

function User() {
  const { userId } = useParams<{ userId: string }>()
  const [user, setUser] = useState<User | null>(null);
  const navigate = useNavigate();
  const updateUser = () => {
    console.log('User:', user);

    if (!user) return;

    axios.put(`${import.meta.env.VITE_API_BASE_URL}/users/${user.id}`, user)
      .then(response => {
        console.log('User updated:', response.data);
        navigate(`/users/${response.data.id}`);
      })
      .catch(error => {
        console.error('Error updating user:', error);
      });
  }

  useEffect(() => {
    const fetchUser = () => {
      axios.get(`${import.meta.env.VITE_API_BASE_URL}/users/${userId}`)
        .then(response => response.data)
        .then(data => setUser(data))
        .catch(error => console.error('Error fetching data:', error))
    }
    fetchUser()
  }, [userId])

  return (
    <>
      <BackButton to={`/users/${userId}`} />
      {user &&
        <Form>
          <Row className="mb-3">
            <Form.Group as={Col} controlId="formGridLogin">
              <Form.Label>Login</Form.Label>
              <Form.Control type="text" placeholder="Enter login" disabled={true} value={user.login} />
            </Form.Group>

            <Form.Group as={Col} controlId="formGridFullName">
              <Form.Label>Full Name</Form.Label>
              <Form.Control type="text" placeholder="Enter full name" value={user.fullName} onChange={(e) => setUser({ ...user, fullName: e.target.value })} />
            </Form.Group>
          </Row>
          <Button variant="primary" onClick={updateUser}>
            Update
          </Button>
        </Form>
      }
    </>
  )
}

export default User
