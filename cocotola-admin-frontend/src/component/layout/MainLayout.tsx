// import { Dialog, Menu, Transition } from '@headlessui/react';
import {useState} from "react";

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

// import logo from '@/assets/react.svg';
// import { useAuth } from '@/lib/auth';
// import { useAuthorization, ROLES } from '@/lib/authorization';
import {
  Box,
  Button,
  Flex,
  MenuContent,
  //   MenuButton,
  //   MenuList,
  MenuItem,
  MenuRoot,
  Spacer,
  Text,
  //   MenuDivider,
} from "@chakra-ui/react";
import { useNavigate } from "react-router";
import { Link, Tabs } from "@chakra-ui/react";
import { Link as ChakraLink } from "@chakra-ui/react";
import { Link as ReactRouterLink } from "react-router";
import {useNavStore} from "@/feature/store/nav";
// import { useAuthStore } from '@/stores/auth';
const NavBar = () => {
  //   const logout = useAuthStore((state) => state.resetTokens);
  const logout = () => {};
  const [currentTab, setCurrentTab] = useState("dashboard");
  let navigate = useNavigate();
  const tab=  useNavStore((state)=>state.tab);
  console.log("tab",tab);

  return (
    <NavBarContainer>
      <Logo />
      {/* <Logo /> */}
      {/* <Spacer /> */}
      {/* <MenuRoot>
        <MenuContent as={Button} colorScheme="pink">
          Profile
        </MenuContent>
        <MenuContent>
          <MenuContent title="Profile">
            <MenuItem value="my-account">My Account</MenuItem>
            <MenuItem value="logout" onClick={() => logout()}>Logout </MenuItem>
          </MenuContent>
          <MenuContent title="Help">
            <MenuItem value="docs">Docs</MenuItem>
            <MenuItem value="faq">FAQ</MenuItem>
          </MenuContent>
        </MenuContent>
      </MenuRoot> */}
      <Tabs.Root defaultValue={tab} variant="subtle">
        <Tabs.List>
          <Tabs.Trigger value="dashboard" asChild>
            <ReactRouterLink to="/tatoeba">Dashboard</ReactRouterLink>

          </Tabs.Trigger>
          <Tabs.Trigger value="tatoeba" asChild>
            <ReactRouterLink to="/tatoeba">Tatoeba</ReactRouterLink>
          </Tabs.Trigger>
        </Tabs.List>
        {/* <Tabs.Content value="members">Manage your team members</Tabs.Content> */}
        {/* <Tabs.Content value="projects">Manage your projects</Tabs.Content> */}
      </Tabs.Root>
      {/* 
      <MenuLinks isOpen={isOpen} /> */}
    </NavBarContainer>
  );
};

const Logo = () => {
  return (
    <Box w="100px" color={"white"}>
      <Text fontSize="lg" fontWeight="bold">
        cocotola
      </Text>
    </Box>
  );
};

type NavBarContainerProps = {
  children: React.ReactNode;
};

const NavBarContainer = ({ children }: NavBarContainerProps) => {
  return (
    <Flex
      as="nav"
      // align="center"
      // justify="space-between"
      // wrap="wrap"
      // w="100%"
      mb={0}
      p={4}
      //   background={'blue.500'}
      // color={'white'}
    >
      {children}
    </Flex>
  );
};
type MainLayoutProps = {
  children: React.ReactNode;
  title: string;
};

export const MainLayout = ({ children }: MainLayoutProps) => {
  return (
    <Box>
      <NavBar />
      <Box>{children}</Box>
    </Box>
  );
};
