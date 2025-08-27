import { NavLink } from "react-router";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faArrowAltCircleLeft } from '@fortawesome/free-regular-svg-icons';
import Button from 'react-bootstrap/Button';

interface BackProps {
  to: string;
}

function BackButton({ to }: BackProps) {
  return (
    <NavLink to={to}>
      <Button variant="outline-secondary"><FontAwesomeIcon icon={faArrowAltCircleLeft} /> Back</Button>
    </NavLink>
  );
}

export default BackButton;
