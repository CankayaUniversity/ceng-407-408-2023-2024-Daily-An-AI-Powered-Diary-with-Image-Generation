import { StatusBar } from 'expo-status-bar';
import { StyleSheet, Text, View, TouchableOpacity, ImageBackground } from 'react-native';
import Header from '../components/Header';

const Home = ({ navigation }: { navigation: any }) => {
  return (
    <Header navigation={navigation} previous="Home" homepage={true}>
      <View style={styles.row1}>
        <TouchableOpacity style={styles.box1} onPress={() => navigation.navigate("Statistics")}>
          <Text style={styles.text}>statistics</Text>
        </TouchableOpacity>
      </View>
      <View style={styles.container}>
        <View style={styles.row2}>
          <TouchableOpacity style={styles.box2} onPress={() => navigation.navigate("YourDaily")}>
            <Text style={styles.text}>your daily</Text>
          </TouchableOpacity>
          <TouchableOpacity style={styles.box2} onPress={() => navigation.navigate("Explore")}>
            <Text style={styles.text}>explore</Text>
          </TouchableOpacity>
        </View>
        <View style={styles.row3}>
          <TouchableOpacity style={styles.box3} onPress={() => navigation.navigate("WriteADaily")}>
            <Text style={styles.text}>write a {'\n'}daily</Text>
          </TouchableOpacity>
          <TouchableOpacity style={styles.box3} onPress={() => navigation.navigate("Profile")}>
            <Text style={styles.text}>profile</Text>
          </TouchableOpacity>
        </View>
        <StatusBar style="auto" />
      </View>
    </Header >
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    flexDirection: 'column',
  },
  row1: {
    height: "25%", 
    paddingStart: 12, 
    paddingEnd: 12, 
    paddingTop: 12, 
  },
  row2: { 
    flexDirection: "row", 
    gap: 12, 
    alignItems: 'center', 
    paddingStart: 12, 
    paddingEnd: 24, 
    marginTop: 12, 
    height: '50%' 
  },
  row3: {
    flexDirection: "row", 
    gap: 12, 
    paddingStart: 12, 
    paddingEnd: 24, 
    height: '25%'
  },
  box1: {
    position: 'relative', 
    width: '100%', 
    height: '100%', 
    justifyContent: 'flex-start', 
    borderRadius: 12, 
    backgroundColor: 'rgba(0, 0, 0, 0.4)',
    elevation: 3,
    shadowOffset: {
      width: 0,
      height: 4
    },
    shadowRadius: 3,
    shadowOpacity: 0.4
  },
  box2: { 
    height: '100%', 
    width: '50%', 
    justifyContent: 'flex-start', 
    borderRadius: 12, 
    backgroundColor: 'rgba(0, 0, 0, 0.4)',
    elevation: 3,
    shadowOffset: {
       width: 0,
       height: 4
    },
    shadowRadius: 3,
    shadowOpacity: 0.4
  },
  box3: { 
    aspectRatio: 1 / 1, 
    width: '50%', 
    marginTop: 12, 
    justifyContent: 
    'flex-start', 
    borderRadius: 12, 
    backgroundColor: 'rgba(0, 0 ,0 , 0.4)',
    elevation: 3,
    shadowOffset: {
      width: 0,
      height: 4
    },
    shadowRadius: 3,
    shadowOpacity: 0.4

  },
  text: {
    textAlign: 'right',
    marginEnd: 10,
    marginTop: 10,
    fontSize: 40,
    fontWeight: '200',
    color: 'white'
  }
});

export default Home;
