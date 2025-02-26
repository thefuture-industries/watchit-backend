import {
	View,
	Text,
	StyleSheet,
	FlatList,
	Image,
	RefreshControl,
	TouchableOpacity,
	ViewToken,
} from 'react-native';

import { SafeAreaView } from 'react-native-safe-area-context';
import { useSharedValue } from 'react-native-reanimated';
import { ChevronRight, Search } from 'lucide-react-native';
import { useState } from 'react';
import { Link, useRouter } from 'expo-router';

import { MovieModel } from '@/types/movie.types';
import { Images } from '@/constants/Images';
import { Colors } from '@/constants/Colors';

import useServerData from '@/hooks/useServerData';
import movieService from '@/services/movie.service';

import Button from '@/components/ui/Button';
import MovieCard from '@/components/ui/MovieCard';

export default function HomeScreen() {
	const viewableItems = useSharedValue<ViewToken[]>([]);

	const {
		data,
		isLoading,
		refetch,
	}: { data: MovieModel[]; isLoading: boolean; refetch: () => any } =
		useServerData(movieService.popular);

	const [refreshing, setRefreshing] = useState<boolean>(false);
	const router = useRouter();

	async function onRefresh() {
		setRefreshing(true);
		await refetch();
		setRefreshing(false);
	}

	return (
		<SafeAreaView className="h-full">
			<View className="my-4 px-4">
				<View className="flex-row justify-between items-center">
					<Link href="/favourite" className="flex-row items-center">
						<Image
							source={Images.user_logo}
							resizeMode="contain"
							className="w-[7vw] h-[7vw] rounded-2xl"
						/>
						<View className="flex-row items-center pl-4">
							<Text className="text-[18px] text-[#fff] font-semibold">
								Гость2781
							</Text>
							<ChevronRight color="#fff" size={20} />
						</View>
					</Link>
					<TouchableOpacity onPress={() => router.push('/search')}>
						<Search color={Colors.app.tint} size={20} />
					</TouchableOpacity>
				</View>
			</View>

			<FlatList
				data={data}
				keyExtractor={(item: MovieModel) => item.id.toString()}
				numColumns={2}
				onViewableItemsChanged={({ viewableItems: vItems }) => {
					viewableItems.value = vItems;
				}}
				renderItem={({ item }) => (
					<View style={{ margin: 8 }}>
						<MovieCard movie={item} viewableItems={viewableItems} />
					</View>
				)}
				ListEmptyComponent={() => (
					<View
						className="flex-column px-5 py-3 bg-[#222] rounded-2xl"
						style={{ margin: 50 }}
					>
						<View className="items-center">
							<Image
								source={Images.empty_data}
								className="w-[55vw] h-[55vw] opacity-[0.7]"
								resizeMode="contain"
							/>
							<Text className="text-[#888] text-[16px] text-center max-w-[75vw] leading-7 mb-4">
								Movies not found please try again later or
								search for movies by filters
							</Text>
						</View>
						{/* <Button text="Movie Filter" /> */}
					</View>
				)}
				refreshControl={
					<RefreshControl
						refreshing={refreshing}
						onRefresh={onRefresh}
					/>
				}
			/>
		</SafeAreaView>
	);
}

const styles = StyleSheet.create({
	text: {
		color: Colors.app.text,
		fontSize: 18,
	},
	input: {
		backgroundColor: '#222',
		borderRadius: 30,
		paddingTop: 17,
		paddingBottom: 17,
		paddingLeft: 25,
		marginTop: 17,
		color: '#fff',
		fontSize: 16,
	},
	quik_genre: {
		backgroundColor: '#222',
		color: '#fff',
		borderRadius: 20,
		padding: 7,
		paddingLeft: 20,
		paddingRight: 20,
		marginRight: 7,
	},
});
