import { StatusBar } from 'expo-status-bar';
import { StyleSheet, Text, View,Alert,TouchableOpacity, ImageBackground } from 'react-native';
import Header from '../components/Header';

const Home = ({navigation}) => {
   return (
      <Header previous="Home" homepage={true}>
      <View style={{height:"25%",paddingStart:10,paddingEnd:10,paddingTop:10 ,borderTopStartRadius:16,borderTopEndRadius:16}}>
      <TouchableOpacity style={{ position:'relative',width:'100%',borderWidth:0.5,borderColor:'gray', height:'100%',opacity:0.85,justifyContent:'flex-start',float:'center',borderRadius:10,backgroundColor:'#0D1326'}} onPress={()=>navigation.navigate("Statistics")}>
          <Text style={styles.text}>statistics</Text>
        </TouchableOpacity>
      </View>
      <View style={styles.container}>
        <View style={{flexDirection:"row",gap:10,alignItems:'center',paddingStart:10,paddingEnd:20,marginTop:10,height:'50%'}}>
        <TouchableOpacity style={{height:'100%', width:'50%',borderWidth:0.5,borderColor:'gray',opacity:0.85,justifyContent:'flex-start',float:'left',borderRadius:10,backgroundColor:'#0D1326'}} onPress={()=>navigation.navigate("YourDaily")}>
          <Text style={styles.text}>your daily</Text>
        </TouchableOpacity>
        <TouchableOpacity style={{ width:'50%', height:'100%',borderWidth:0.5,borderColor:'gray',opacity:0.85,justifyContent:'flex-start',float:'left',borderRadius:10,backgroundColor:'#0D1326'}} onPress={()=>navigation.navigate("Explore")}>
          <Text style={styles.text}>explore</Text>
        </TouchableOpacity>
        </View>
        <View style={{flexDirection:"row",gap:10,paddingStart:10,paddingEnd:20,height:'25%'}}>
        <TouchableOpacity style={{aspectRatio:1/1, width:'50%',borderWidth:0.5,borderColor:'gray',opacity:0.85, marginTop: 10,justifyContent:'flex-start',float:'left',borderRadius:10,backgroundColor:'#0D1326'}} onPress={()=>navigation.navigate("WriteADaily")}>
          <Text style={styles.text}>write a {'\n'}daily</Text>
        </TouchableOpacity>
        <TouchableOpacity style={{aspectRatio:1/1, width:'50%',borderWidth:0.5,borderColor:'gray',opacity:0.85, marginTop: 10,justifyContent:'flex-start',float:'left',borderRadius:10,backgroundColor:'#0D1326'}} onPress={()=>navigation.navigate("Profile")}>
          <Text style={styles.text}>profile</Text>
        </TouchableOpacity>
        </View>
        <StatusBar style="auto" />
      </View>
      </Header>
   );
}

const styles = StyleSheet.create({
   container: {
     flex: 1,
     flexDirection:'column',
   },
   text: {
     textAlign:'right',
     marginEnd:10,
     marginTop:10,
     fontSize:40,
     fontWeight:'200',
     color:'white'
   }
 });

export default Home;
