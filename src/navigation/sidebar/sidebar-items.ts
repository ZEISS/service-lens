import {
  AirVent,
  DraftingCompass,
  Dumbbell,
  Glasses,
  LayoutDashboard,
  type LucideIcon,
  PencilRuler,
  Settings,
  Tags,
  UserCog,
  Users,
  Wrench,
} from "lucide-react"

export interface NavSubItem {
  title: string
  url: string
  icon?: LucideIcon
  comingSoon?: boolean
  newTab?: boolean
  isNew?: boolean
}

export interface NavMainItem {
  title: string
  url: string
  icon?: LucideIcon
  subItems?: NavSubItem[]
  comingSoon?: boolean
  newTab?: boolean
  isNew?: boolean
}

export interface NavGroup {
  id: number
  label?: string
  items: NavMainItem[]
}

export const sidebarItems: NavGroup[] = [
  {
    id: 1,
    label: "Design",
    items: [
      {
        title: "Default",
        url: "/dashboard/default",
        icon: LayoutDashboard,
      },
      {
        title: "Designs",
        url: "/designs",
        icon: PencilRuler,
      },
    ],
  },
  {
    id: 2,
    label: "Review",
    items: [
      {
        title: "Workloads",
        url: "/workloads",
        icon: Dumbbell,
        comingSoon: false,
      },
      {
        title: "Lenses",
        url: "/lenses",
        icon: Glasses,
        comingSoon: false,
      },
      {
        title: "Environments",
        url: "/environments",
        icon: AirVent,
        comingSoon: false,
      },
      {
        title: "Profiles",
        url: "/profiles",
        icon: DraftingCompass,
        comingSoon: false,
      },
    ],
  },
  {
    id: 3,
    label: "Misc",
    items: [
      {
        title: "Tags",
        url: "/tags",
        icon: Tags,
      },
      {
        title: "Settings",
        url: "/settings/profile",
        icon: Settings,
        subItems: [
          {
            title: "Profile",
            url: "/settings/profile",
            icon: UserCog,
            newTab: false,
          },
          {
            title: "Account",
            url: "/settings/account",
            icon: Wrench,
            newTab: false,
          },
        ],
      },
    ],
  },
  {
    id: 4,
    label: "Admin",
    items: [
      {
        title: "Users",
        url: "/admin/users",
        icon: Users,
      },
      {
        title: "Settings",
        url: "/admin/settings",
        icon: Settings,
      },
    ],
  },
]
