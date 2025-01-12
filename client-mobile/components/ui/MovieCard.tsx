import {
  View,
  Text,
  Image,
  Platform,
  StyleSheet,
  TouchableOpacity,
  ViewToken,
} from 'react-native';
import React, { useRef, useState } from 'react';
import { MovieModel } from '@/types/movie.types';
import configService from '@/services/config.service';
import { Images } from '@/constants/Images';
import { Star } from 'lucide-react-native';

import Animated, {
  useAnimatedStyle,
  withTiming,
} from 'react-native-reanimated';

interface MovieCardProps {
  movie: MovieModel;
  viewableItems: Animated.SharedValue<ViewToken[]>;
}

export default function MovieCard({ movie, viewableItems }: MovieCardProps) {
  // **
  // Platform
  const platform = Platform.OS;
  let src = '';

  if (platform === 'ios' || platform === 'android') {
    src = `${configService.ReturnConfig().SERVER_URL_MOBILE}/image/w500${movie.poster_path}`;
  } else {
    src = `${configService.ReturnConfig().SERVER_URL_WEB}/image/w500${movie.poster_path}`;
  }

  // **
  // Image
  const imgRef = useRef(null);
  const [poster, setPoster] = useState<string>('');
  const [loaded, setLoaded] = useState<boolean>(true);

  const handleLoaded = () => {
    setPoster(src);
    setLoaded(false);
  };

  const handleError = () => {
    setPoster(Images.default_image);
    setLoaded(false);
  };

  const rStyle = useAnimatedStyle(() => {
    const isVisible = Boolean(
      viewableItems.value
        .filter((item) => item.isViewable)
        .find((viewableItem) => viewableItem.item.id === movie.id)
    );

    return {
      opacity: withTiming(isVisible ? 1 : 0),
      transform: [
        {
          scale: withTiming(isVisible ? 1 : 0.6),
        },
      ],
    };
  }, []);

  const ListMovie = React.memo(({ movie }: { movie: MovieModel }) => {
    return (
      <Animated.View style={rStyle}>
        <View ref={imgRef}>
          {!loaded && poster ? (
            <Image
              source={typeof poster === 'number' ? poster : { uri: poster }}
              style={{ width: 180, height: 250, borderRadius: 10 }}
            />
          ) : (
            <Image
              source={{ uri: src }}
              style={{ width: 180, height: 250, borderRadius: 10 }}
              onLoad={handleLoaded}
              onError={handleError}
            />
          )}
        </View>
        <View style={{ marginTop: 12 }}>
          <Text style={styles.text_date}>
            {movie.release_date.split('-')[0]}
          </Text>
          <Text style={styles.text_title}>{movie.title}</Text>
          <View className="flex-row items-center" style={{ marginTop: 3 }}>
            <Star fill="#ffdb3d" size={15} />
            <Text style={styles.text_average}>{movie.vote_average}</Text>
          </View>
        </View>
      </Animated.View>
    );
  });

  return (
    <TouchableOpacity onPress={() => console.log('click')}>
      <ListMovie movie={movie} />
    </TouchableOpacity>
  );
}

const styles = StyleSheet.create({
  text_title: {
    color: '#fff',
    maxWidth: 150,
    fontWeight: 700,
    fontSize: 16,
    marginTop: 3,
  },
  text_date: {
    color: '#888',
    maxWidth: 150,
    fontWeight: 400,
    fontSize: 16,
  },
  text_average: {
    color: '#999',
    maxWidth: 150,
    fontWeight: 400,
    fontSize: 15,
    marginLeft: 3,
  },
});
