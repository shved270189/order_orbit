import { useState, useEffect } from 'react'
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import { NavLink } from "react-router";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faEdit, faTrashCan, faEye } from '@fortawesome/free-regular-svg-icons';
import Moment from 'moment';

type User = {
  id: number;
  fullName: string;
  login: string;
  updatedAt: string;
  createdAt: string;
};

function Users() {
  const [users, setUsers] = useState<User[]>([])

  useEffect(() => {
    const fetchData = () => {
      fetch(`http://localhost:${import.meta.env.VITE_PORT}/users`)
        .then(response => response.text())
        .then(data => setUsers(JSON.parse(data)))
        .catch(error => console.error('Error fetching data:', error))
    }
    fetchData()
  }, [])

  return (
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
              <th scope="col">Full Name</th>
              <th scope="col">Login</th>
              <th scope="col">Updated At</th>
              <th scope="col">Created At</th>
              <th scope="col">Actions</th>
            </tr>
          </thead>
          <tbody>
            {users.map((user) => {
              return (
                <tr key={user.id}>
                  <th scope="row">{user.id}</th>
                  <td>{user.fullName}</td>
                  <td>{user.login}</td>
                  <td>{Moment(user.updatedAt).format('MMMM Do YYYY, h:mm:ss a')}</td>
                  <td>{Moment(user.createdAt).format('MMMM Do YYYY, h:mm:ss a')}</td>
                  <td>
                    <NavLink className="link" to={`/user/${user.id}`}>
                      <FontAwesomeIcon icon={faEye} />
                    </NavLink>
                    <NavLink className="link" to={`/user/${user.id}`}>
                      <FontAwesomeIcon icon={faEdit} />
                    </NavLink>
                    <NavLink className="link" to={`/user/${user.id}`}>
                      <FontAwesomeIcon icon={faTrashCan} />
                    </NavLink>
                  </td>
                </tr>
              )
            })}
          </tbody>
        </table>
      )}
      </Col>
    </Row>
  )
}

export default Users
