import { Tabs } from 'expo-router';
import MaterialIcons from '@expo/vector-icons/MaterialIcons';
import { Colors } from '@/constants/Colors';

// You can explore the built-in icon families and icons on the web at https://icons.expo.fyi/
function TabBarIcon(props: {
  name: React.ComponentProps<typeof MaterialIcons>['name'];
  color: string;
}) {
  return <MaterialIcons size={26} style={{ marginBottom: -3 }} {...props} />;
}

export default function TabLayout() {
  return (
    <Tabs
      screenOptions={{
        tabBarStyle: {
          backgroundColor: Colors.app.gray_background,
          borderColor: Colors.app.border,
        },
      }}
    >
      <Tabs.Screen
        name="index"
        options={{
          title: 'Home',
          headerShown: false,
          tabBarIcon: ({ color }) => <TabBarIcon name="home" color={color} />,
        }}
      />
      <Tabs.Screen
        name="favourite"
        options={{
          title: 'Favourite',
          headerShown: false,
          tabBarIcon: ({ color }) => (
            <TabBarIcon name="favorite" color={color} />
          ),
        }}
      />
      <Tabs.Screen
        name="video"
        options={{
          title: 'Video',
          headerShown: false,
          tabBarIcon: ({ color }) => (
            <TabBarIcon name="video-library" color={color} />
          ),
        }}
      />
      <Tabs.Screen
        name="movie"
        options={{
          title: 'Movie',
          headerShown: false,
          tabBarIcon: ({ color }) => <TabBarIcon name="movie" color={color} />,
        }}
      />
      <Tabs.Screen
        name="story"
        options={{
          title: 'Story',
          headerShown: false,
          tabBarIcon: ({ color }) => (
            <TabBarIcon name="short-text" color={color} />
          ),
        }}
      />
    </Tabs>
  );
}
