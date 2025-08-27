import { useState, useEffect } from 'react'
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Moment from 'moment';
import {useParams} from "react-router";
import axios from 'axios';
import BackButton from '../common/BackButton';
import { NavLink } from "react-router";
import { Button } from 'react-bootstrap';
import { useNavigate } from "react-router";
import Modal from 'react-bootstrap/Modal';

type User = {
  id: number;
  fullName: string;
  login: string;
  updatedAt: string;
  createdAt: string;
};

function User() {
  const { userId } = useParams<{ userId: string }>();
  const navigate = useNavigate();
  const [user, setUser] = useState<User | null>(null);
  const [showDeletePrompt, setShowDeletePrompt] = useState(false);
  const deleteUser = () => {
    axios.delete(`${import.meta.env.VITE_API_BASE_URL}/users/${userId}`)
      .then(() => {
        setShowDeletePrompt(false);
        navigate('/users');
      })
      .catch(error => console.error('Error deleting user:', error));
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
    <Row>
      <Col>
        <BackButton to="/users" />
      </Col>
      <Col className="d-grid gap-2 d-md-flex justify-content-md-end">
        <NavLink to={`/users/${userId}/edit`}>
          <Button variant="primary">Edit</Button>
        </NavLink>
        <Button variant="danger" onClick={() => setShowDeletePrompt(true)}>Delete</Button>
      </Col>
    </Row>
    <Row>
      <Col>
      {user === null && (
        <div>
          <h2>Loading...</h2>
        </div>
      )}
      {user && (
        <div>
          <h2>{user.fullName}</h2>
          <p>Login: {user.login}</p>
          <p>Updated At: {Moment(user.updatedAt).format('MMMM Do YYYY, h:mm:ss a')}</p>
          <p>Created At: {Moment(user.createdAt).format('MMMM Do YYYY, h:mm:ss a')}</p>
        </div>
      )}
      </Col>
    </Row>
    {showDeletePrompt && user && (
      <Modal show={showDeletePrompt} onHide={() => setShowDeletePrompt(false)}>
        <Modal.Header closeButton>
          <Modal.Title>Confirm Deletion</Modal.Title>
        </Modal.Header>
        <Modal.Body>Are you sure you want to delete user "{user.fullName} ({user.login})"?</Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={() => setShowDeletePrompt(false)}>
            Close
          </Button>
          <Button variant="danger" onClick={() => deleteUser()}>
            Delete User
          </Button>
        </Modal.Footer>
      </Modal>
    )}
    </>
  )
}

export default User
