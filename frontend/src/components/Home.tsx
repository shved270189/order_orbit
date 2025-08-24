import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Image from 'react-bootstrap/Image';

function Home() {
  return (
    <Row>
      <Col>
        <h2>Hello! Nice to meet you in Order Orbit!</h2>
        <Image src="/logo.jpeg" roundedCircle />
      </Col>
    </Row>
  )
}
export default Home
