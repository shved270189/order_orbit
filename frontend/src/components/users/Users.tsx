import { useState, useEffect } from 'react'
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import { NavLink } from "react-router";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faEdit, faTrashCan, faEye } from '@fortawesome/free-regular-svg-icons';
import Moment from 'moment';
import axios from 'axios';
import { Button } from 'react-bootstrap';
import Modal from 'react-bootstrap/Modal';

type User = {
  id: number;
  fullName: string;
  login: string;
  updatedAt: string;
  createdAt: string;
};

function Users() {
  const [users, setUsers] = useState<User[]>([])
  const [userForDelete, setUserForDelete] = useState<User | null>(null);
  const fetchData = () => {
    axios.get(`${import.meta.env.VITE_API_BASE_URL}/users`)
      .then(response => response.data)
      .then(data => setUsers(data))
      .catch(error => console.error('Error fetching data:', error))
  }
  const deleteUser = (user: User) => {
    axios.delete(`${import.meta.env.VITE_API_BASE_URL}/users/${user.id}`)
      .then(() => {
        setUsers(users.filter(u => u.id !== user.id));
        setUserForDelete(null);
        fetchData();
      })
      .catch(error => console.error('Error deleting user:', error));
  }

  useEffect(() => {
    fetchData()
  }, [])

  return (
    <>
      <Row>
        <Col>
          <h2>Users</h2>
        </Col>
        <Col className="d-grid gap-2 d-md-flex justify-content-md-end">
          <NavLink to={`/users/new`}>
            <Button variant="primary">Add New</Button>
          </NavLink>
        </Col>
      </Row>
      <Row>
        <Col>
        {users.length === 0 && (
          <div>
            <h2>Loading...</h2>
          </div>
        )}
        {users.length > 0 && (
          <table className="table table-striped">
            <thead>
              <tr>
                <th scope="col">ID</th>
                <th scope="col">Login</th>
                <th scope="col">Full Name</th>
                <th scope="col">Updated At</th>
                <th scope="col">Created At</th>
                <th scope="col">Actions</th>
              </tr>
            </thead>
            <tbody>
              {users.map((user) => {
                return (
                  <tr key={user.id}>
                    <th scope="row">
                      <NavLink className="link" to={`/users/${user.id}`}>{user.id}</NavLink>
                    </th>
                    <td>{user.login}</td>
                    <td>{user.fullName}</td>
                    <td>{Moment(user.updatedAt).format('MMMM Do YYYY, h:mm:ss a')}</td>
                    <td>{Moment(user.createdAt).format('MMMM Do YYYY, h:mm:ss a')}</td>
                    <td>
                      <NavLink to={`/users/${user.id}`}>
                        <Button variant="link">
                          <FontAwesomeIcon icon={faEye} />
                        </Button>
                      </NavLink>
                      <NavLink to={`/users/${user.id}/edit`}>
                        <Button variant="link">
                          <FontAwesomeIcon icon={faEdit} />
                        </Button>
                      </NavLink>
                      <Button variant="link" onClick={() => setUserForDelete(user)}>
                        <FontAwesomeIcon icon={faTrashCan} />
                      </Button>
                    </td>
                  </tr>
                )
              })}
            </tbody>
          </table>
        )}
        </Col>
      </Row>
      {userForDelete && (
        <Modal show={!!userForDelete} onHide={() => setUserForDelete(null)}>
          <Modal.Header closeButton>
            <Modal.Title>Confirm Deletion</Modal.Title>
          </Modal.Header>
          <Modal.Body>Are you sure you want to delete user "{userForDelete.fullName} ({userForDelete.login})"?</Modal.Body>
          <Modal.Footer>
            <Button variant="secondary" onClick={() => setUserForDelete(null)}>
              Close
            </Button>
            <Button variant="danger" onClick={() => deleteUser(userForDelete)}>
              Delete User
            </Button>
          </Modal.Footer>
        </Modal>
      )}
    </>
  )
}

export default Users
