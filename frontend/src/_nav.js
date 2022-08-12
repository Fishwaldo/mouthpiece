export default [
  {
    component: 'CNavItem',
    name: 'Dashboard',
    to: '/dashboard',
    icon: 'cil-speedometer',
    badge: {
      color: 'primary',
      text: 'NEW',
    },
  },
  {
    component: 'CNavTitle',
    name: 'Apps',
  },
  {
    component: 'CNavItem',
    name: 'Colors',
    to: '/apps/',
    icon: 'cil-puzzle',
  },
  {
    component: 'CNavTitle',
    name: 'Users',
  },
  {
    component: 'CNavItem',
    name: 'Users',
    to: '/users',
    icon: 'cil-puzzle',
  },
  {
    component: 'CNavItem',
    name: 'Transports',
    to: '/transports',
    icon: 'cil-cursor',
  },
  {
    component: 'CNavGroup',
    name: 'Notifications',
    to: '/notifications',
    icon: 'cil-bell',
    items: [
      {
        component: 'CNavItem',
        name: 'Alerts',
        to: '/notifications/alerts',
      },
      {
        component: 'CNavItem',
        name: 'Badges',
        to: '/notifications/badges',
      },
      {
        component: 'CNavItem',
        name: 'Modals',
        to: '/notifications/modals',
      },
    ],
  },
  {
    component: 'CNavTitle',
    name: 'Settings',
  },
  {
    component: 'CNavGroup',
    name: 'Settings',
    to: '/settings',
    icon: 'cil-star',
  },
]
