import { StatusBar } from 'expo-status-bar';
import { StyleSheet, Text, View,Alert,TouchableOpacity, ImageBackground } from 'react-native';
import Header from '../components/Header';

const Profile = ({navigation}:{navigation:any}) => {
   return (
      <Header navigation={navigation} previous="Home" homepage={true}>
            <View style={styles.container}>
               <View style={styles.columns}>
                  <View style={styles.rows}>
                     <View style={{width:'40%', height:'40%', paddingTop:12,borderRadius:12, alignItems: 'center',backgroundColor:'rgba(0,0,0,0.4)'}}>
                        <View style={styles.avatarPlaceholder}/>
                     </View>
                     <View style={{width:'60%',height:'40%', paddingTop:12,borderRadius:12,backgroundColor:'rgba(0,0,0,0.4)'}}>
                        <Text style={styles.name}>daily dev</Text>
                        <Text style={styles.username}>@daily_dev</Text>
                     </View>
                  </View>
                  <View style={styles.rows}>
                     <View style={{width:'40%', height:'20%', paddingTop:12,borderRadius:12,backgroundColor:'rgba(0,0,0,0.4)'}}>
                        <Text style={styles.username}>test</Text>
                     </View> 
                  </View>
                  
               </View>
            </View>
      </Header>
   );
}

const styles = StyleSheet.create({
   container: {
      flex: 1,
      flexDirection: 'column',
      justifyContent: 'flex-start',
      padding:12
   },
   columns: {
      flex: 1,
      flexDirection: 'column',
      justifyContent: 'flex-start',
      borderRadius:12,
      padding:12,
      backgroundColor:'rgba(0,0,0,0.4)'
   },
   rows: {
      flex: 1,
      flexDirection: 'row',
      justifyContent: 'flex-start',
      borderRadius:12,
   },
   box: {
      paddingStart:12,
      paddingEnd:12,
      paddingTop:12, 
      borderTopStartRadius:12,
      borderTopEndRadius:12
   },
   name: {
      textAlign:'right',
      marginEnd:12.5,
      fontSize:50,
      fontWeight:'500',
      color:'white',
      opacity: 1.0
   },
   username: {
      textAlign:'right',
      marginEnd:12.5,
      fontSize:20,
      fontWeight:'200',
      color:'white',
      opacity: 0.5
   },
   avatar: {
      width: 128,
      height: 128,
      borderRadius: 50,
   },
   avatarPlaceholder: {
      width: '80%',
      height: '80%',
      borderRadius: 100,
      backgroundColor: '#FFFFFF',
      opacity: 0.5,
      alignSelf: 'center',
      justifyContent: 'center',
   }
   
});

export default Profile;
