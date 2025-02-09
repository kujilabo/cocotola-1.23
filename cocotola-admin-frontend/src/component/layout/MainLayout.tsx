// import { Dialog, Menu, Transition } from '@headlessui/react';
import { useState } from "react";

// import { Dialog, Transition } from '@headlessui/react';
// import {
//   UserIcon,
//   FolderIcon,
//   HomeIcon,
//   // MenuAlt2Icon,
//   // UsersIcon,
//   // XIcon,
// } from '@heroicons/react/24/outline';
// import { clsx } from 'clsx';
// import { NavLink, Link } from 'react-router-dom';

import { useNavStore } from "@/feature/store/nav";
import AdbIcon from "@mui/icons-material/Adb";
import MenuIcon from "@mui/icons-material/Menu";
import AppBar from "@mui/material/AppBar";
import Avatar from "@mui/material/Avatar";
import Box from "@mui/material/Box";
import Breadcrumbs from "@mui/material/Breadcrumbs";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import IconButton from "@mui/material/IconButton";
import Link from "@mui/material/Link";
import Menu from "@mui/material/Menu";
import MenuItem from "@mui/material/MenuItem";
import Toolbar from "@mui/material/Toolbar";
import Tooltip from "@mui/material/Tooltip";
import Typography from "@mui/material/Typography";
// import logo from '@/assets/react.svg';
// import { useAuth } from '@/lib/auth';
// import { useAuthorization, ROLES } from '@/lib/authorization';
// import { Link as ChakraLink } from "@chakra-ui/react";
import { useNavigate } from "react-router";
import { Link as ReactRouterLink } from "react-router";

const pages = ["dashboard", "tatoeba", "Blog"];
const settings = ["Profile", "Account", "Dashboard", "Logout"];
// import { useAuthStore } from '@/stores/auth';
const NavBar = () => {
  const navigate = useNavigate();
  const [anchorElNav, setAnchorElNav] = useState<null | HTMLElement>(null);
  const [anchorElUser, setAnchorElUser] = useState<null | HTMLElement>(null);

  const handleOpenNavMenu = (event: React.MouseEvent<HTMLElement>) => {
    console.log("handleOpenNavMenu");
    setAnchorElNav(event.currentTarget);
    // navigate("/" + event.currentTarget);
  };
  const handleOpenUserMenu = (event: React.MouseEvent<HTMLElement>) => {
    console.log("handleOpenUserMenu");
    setAnchorElUser(event.currentTarget);
  };

  const handleCloseNavMenu = (page: string) => {
    setAnchorElNav(null);
    navigate(`/${page}`);
  };

  const handleCloseUserMenu = () => {
    setAnchorElUser(null);
  };
  //   const logout = useAuthStore((state) => state.resetTokens);
  const logout = () => {};
  const [currentTab, setCurrentTab] = useState("dashboard");
  const tab = useNavStore((state) => state.tab);
  // console.log("tab", tab);
  return (
    <AppBar position="static">
      <Container maxWidth="xl">
        <Toolbar disableGutters>
          <AdbIcon sx={{ display: { xs: "none", md: "flex" }, mr: 1 }} />
          <Typography
            variant="h6"
            noWrap
            component="a"
            href="#app-bar-with-responsive-menu"
            sx={{
              mr: 2,
              display: { xs: "none", md: "flex" },
              fontFamily: "monospace",
              fontWeight: 700,
              letterSpacing: ".3rem",
              color: "inherit",
              textDecoration: "none",
            }}
          >
            LOGO
          </Typography>

          <Box sx={{ flexGrow: 1, display: { xs: "flex", md: "none" } }}>
            <IconButton
              size="large"
              aria-label="account of current user"
              aria-controls="menu-appbar"
              aria-haspopup="true"
              onClick={handleOpenNavMenu}
              color="inherit"
            >
              <MenuIcon />
            </IconButton>
            <Menu
              id="menu-appbar"
              anchorEl={anchorElNav}
              anchorOrigin={{
                vertical: "bottom",
                horizontal: "left",
              }}
              keepMounted
              transformOrigin={{
                vertical: "top",
                horizontal: "left",
              }}
              open={Boolean(anchorElNav)}
              onClose={handleCloseNavMenu}
              sx={{ display: { xs: "block", md: "none" } }}
            >
              {pages.map((page) => (
                <MenuItem key={page} onClick={() => handleCloseNavMenu(page)}>
                  <Typography sx={{ textAlign: "center" }}>{page}</Typography>
                </MenuItem>
              ))}
            </Menu>
          </Box>
          <AdbIcon sx={{ display: { xs: "flex", md: "none" }, mr: 1 }} />
          <Typography
            variant="h5"
            noWrap
            component="a"
            href="#app-bar-with-responsive-menu"
            sx={{
              mr: 2,
              display: { xs: "flex", md: "none" },
              flexGrow: 1,
              fontFamily: "monospace",
              fontWeight: 700,
              letterSpacing: ".3rem",
              color: "inherit",
              textDecoration: "none",
            }}
          >
            LOGO
          </Typography>
          <Box sx={{ flexGrow: 1, display: { xs: "none", md: "flex" } }}>
            {pages.map((page) => (
              <Button
                key={page}
                onClick={() => handleCloseNavMenu(page)}
                sx={{ my: 2, color: "white", display: "block" }}
              >
                {page}
              </Button>
            ))}
          </Box>
          <Box sx={{ flexGrow: 0 }}>
            <Tooltip title="Open settings">
              <IconButton onClick={handleOpenUserMenu} sx={{ p: 0 }}>
                <Avatar alt="Remy Sharp" src="/static/images/avatar/2.jpg" />
              </IconButton>
            </Tooltip>
            <Menu
              sx={{ mt: "45px" }}
              id="menu-appbar"
              anchorEl={anchorElUser}
              anchorOrigin={{
                vertical: "top",
                horizontal: "right",
              }}
              keepMounted
              transformOrigin={{
                vertical: "top",
                horizontal: "right",
              }}
              open={Boolean(anchorElUser)}
              onClose={handleCloseUserMenu}
            >
              {settings.map((setting) => (
                <MenuItem key={setting} onClick={handleCloseUserMenu}>
                  <Typography sx={{ textAlign: "center" }}>
                    {setting}
                  </Typography>
                </MenuItem>
              ))}
            </Menu>
          </Box>
        </Toolbar>
      </Container>
    </AppBar>
  );
  // return (
  //   <NavBarContainer>
  //     <Logo />
  //     {/* <Logo /> */}
  //     {/* <Spacer /> */}
  //     {/* <MenuRoot>
  //       <MenuContent as={Button} colorScheme="pink">
  //         Profile
  //       </MenuContent>
  //       <MenuContent>
  //         <MenuContent title="Profile">
  //           <MenuItem value="my-account">My Account</MenuItem>
  //           <MenuItem value="logout" onClick={() => logout()}>Logout </MenuItem>
  //         </MenuContent>
  //         <MenuContent title="Help">
  //           <MenuItem value="docs">Docs</MenuItem>
  //           <MenuItem value="faq">FAQ</MenuItem>
  //         </MenuContent>
  //       </MenuContent>
  //     </MenuRoot> */}
  //     <Tabs.Root defaultValue={tab} variant="subtle">
  //       <Tabs.List>
  //         <Tabs.Trigger value="dashboard" asChild>
  //           <ReactRouterLink to="/tatoeba">Dashboard</ReactRouterLink>
  //         </Tabs.Trigger>
  //         <Tabs.Trigger value="tatoeba" asChild>
  //           <ReactRouterLink to="/tatoeba">Tatoeba</ReactRouterLink>
  //         </Tabs.Trigger>
  //       </Tabs.List>
  //       {/* <Tabs.Content value="members">Manage your team members</Tabs.Content> */}
  //       {/* <Tabs.Content value="projects">Manage your projects</Tabs.Content> */}
  //     </Tabs.Root>
  //     {/*
  //     <MenuLinks isOpen={isOpen} /> */}
  //   </NavBarContainer>
  // );
};

// const Logo = () => {
//   return (
//     <Box w="100px" color={"white"}>
//       <Text fontSize="lg" fontWeight="bold">
//         cocotola
//       </Text>
//     </Box>
//   );
// };

// type NavBarContainerProps = {
//   children: React.ReactNode;
// };

// const NavBarContainer = ({ children }: NavBarContainerProps) => {
//   return (
//     <Flex
//       as="nav"
//       // align="center"
//       // justify="space-between"
//       // wrap="wrap"
//       // w="100%"
//       mb={0}
//       p={4}
//       //   background={'blue.500'}
//       // color={'white'}
//     >
//       {children}
//     </Flex>
//   );
// };
type MainLayoutProps = {
  children: React.ReactNode;
  title: string;
};

export const MainLayout = ({ children }: MainLayoutProps) => {
  return (
    <Box>
      <NavBar />
      <Breadcrumbs aria-label="breadcrumb">
        <Link underline="hover" color="inherit" href="/">
          MUI
        </Link>
        <Link
          underline="hover"
          color="inherit"
          href="/material-ui/getting-started/installation/"
        >
          Core
        </Link>
        <Typography sx={{ color: "text.primary" }}>Breadcrumbs</Typography>
      </Breadcrumbs>
      <Box>{children}</Box>
    </Box>
  );
};
