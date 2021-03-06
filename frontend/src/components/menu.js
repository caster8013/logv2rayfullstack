import { useSelector, useDispatch } from "react-redux";
import { Container, Navbar, Nav, Button } from "react-bootstrap";
import { logout } from "../store/login";
import { alert, success } from "../store/message";
import AddUser from "./adduser";
import AddNode from "./addNode";
import axios from "axios";

const Menu = () => {
	const loginState = useSelector((state) => state.login);
	const dispatch = useDispatch();

	const handleLogout = (e) => {
		dispatch(logout());
	};
	const handleWriteToDB = (e) => {
		axios({
			method: "get",
			url: process.env.REACT_APP_API_HOST + "writetodb",
			headers: { token: loginState.token },
		})
			.then((response) => {
				dispatch(success({ show: true, content: response.data.message }));
			})
			.catch((err) => {
				dispatch(alert({ show: true, content: err.toString() }));
			});
	};

	return (
		<Navbar bg="dark" variant="dark" expand="lg">
			<Container>
				<Navbar.Brand href="/">Logv2ray Frontend</Navbar.Brand>
				<Navbar.Toggle aria-controls="basic-navbar-nav" />
				<Navbar.Collapse id="basic-navbar-nav">
					{loginState.jwt.Role === "admin" && (
						<Nav>
							<Nav.Link href="/home">User Management Panel</Nav.Link>
						</Nav>
					)}
					<Nav>
						<Nav.Link href="/mypanel">My Panel</Nav.Link>
					</Nav>
					<Nav>
						<Nav.Link href="/macos">MacOS</Nav.Link>
					</Nav>
					<Nav>
						<Nav.Link href="/windows">Windows</Nav.Link>
					</Nav>
					<Nav>
						<Nav.Link href="/iphone">IPhone/IPad</Nav.Link>
					</Nav>
					<Nav>
						<Nav.Link href="/android">Android</Nav.Link>
					</Nav>
				</Navbar.Collapse>
				<Navbar.Collapse className="justify-content-end">
					{loginState.jwt.Role === "admin" && (
						<>
							<Nav>
								<Button
									variant="primary"
									size="sm"
									onClick={handleWriteToDB}
								>
									WriteToDB
								</Button>
							</Nav>
							<Nav className="m-1">
									<AddNode btnName="AddNode" />
							</Nav>
							<Nav className="">
									<AddUser btnName="AddUser" />
							</Nav>
						</>
					)}
					<Nav>
						<Navbar.Text className="mx-2">
							Signed in as: <b>{loginState.jwt.Email}</b>
						</Navbar.Text>
					</Nav>
					<Nav>
							<Button
								variant="outline-secondary"
								size="sm"
								onClick={handleLogout}
							>
								logout
							</Button>
					</Nav>
				</Navbar.Collapse>
			</Container>
		</Navbar>
	);
};

export default Menu;
