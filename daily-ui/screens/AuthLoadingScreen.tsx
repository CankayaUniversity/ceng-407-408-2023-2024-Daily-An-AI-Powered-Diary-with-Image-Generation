import React, { useEffect } from 'react';
import { View } from 'react-native';
import { UserToken, storageKeys, storageUtils } from '../libs';
import Header from '../components/Header';

const AuthLoadingScreen = ({ navigation }: { navigation: any }) => {
  useEffect(() => {
    storageUtils.getItem<UserToken>(storageKeys.bearerToken).then((token) => {
      navigation.navigate(token?.token ? 'Home' : 'Login');
    });
  }, []);

  return (
    <Header navigation={navigation} previous="AuthLoadingScreen" homepage={false}>
      <View style={{ flex: 1, justifyContent: 'center', alignItems: 'center' }}>
      </View>
    </Header>
  );
};

export default AuthLoadingScreen;
