import { StatusBar } from 'expo-status-bar';
import { StyleSheet, Text, View,Alert,TouchableOpacity, ImageBackground,ScrollView } from 'react-native';
import Header from '../components/Header';
import { useGetBadges } from '../libs';
import { Ionicons, FontAwesome } from '@expo/vector-icons';

const Profile = ({navigation}:{navigation:any}) => {

   const {data,isLoading} = useGetBadges()
   const badgeDict = {
      "Beginner Writer":["pencil","white","Beginner Writer\n\nYou wrote your first daily!"],
      "Prolific Writer":["pencil","pink","Prolific Writer\n\nYou're on fire! 100 written daily!"],
      "Master Writer":["pencil","purple","Master Writer\n\nIncredible! You've reached 1000 daily!"],
      "One Week Obsessed":["calendar","purple","One Week Obsessed\n\nYou've a streak for one whole week!"],
      "Admired":["heart","white","Admired\n\nYour work received its first like!"],
      "Liked by Many":["heart","pink","Liked by Many\n\nYou've garnered 100 likes from your readers!"],
      "Influence":["heart","purple","Influencer\n\nA thousand like!"],
      "They Look Here":["eye","pink","They Look Here\n\nYour writing caught someone's eye!"],
      "Popular Author":["eye","purple","Popular Author\n\nYou've achieved 1000 views!"],

   }
   return (
      <Header navigation={navigation} previous="Home" homepage={false}>
         <View style={styles.container}>
            <View style={styles.container2}>
         <ScrollView showsVerticalScrollIndicator={false} contentContainerStyle={{ flexDirection: "column", flexWrap: 'nowrap', justifyContent: 'flex-start', gap: 14, alignItems: 'center', marginTop: 10, paddingBottom: 90 }}>
         {
            (isLoading ) &&
            <Text style={{ alignItems: 'center', justifyContent: 'center', fontSize: 40, color: 'white'}}>Loading</Text>
         }
         {!isLoading && data?.map((el, index) => {
            return (
               <View
               key={index}
               style={{
               height: '40%',
               width: '90%',
               borderWidth: 0.5,
               borderColor: 'gray',
               alignItems: "center",
               justifyContent: "center",
               borderRadius: 10,
               backgroundColor:'black',
               opacity:0.6
               }}
            >
               <View style={{flexDirection:'row' ,width: '100%', height: '100%', alignItems: 'center', justifyContent: 'flex-start',gap:18,paddingStart:10}}>
               <FontAwesome name={badgeDict[el][0]} color={badgeDict[el][1]} size={32} style={{shadowOffset:{width:0,height:0},shadowRadius:2,shadowColor:badgeDict[el][1],shadowOpacity:1}}></FontAwesome>
               <Text style={{ color: "white", fontSize: 20, fontWeight: '200'}}>{badgeDict[el][2]}</Text>
               </View>
            </View>
            );
         })
         }
         </ScrollView>
         </View>
         </View>
      </Header>
   );
}

const styles = StyleSheet.create({
   container: {
      height: '90%',
      padding:8
   },
   container2: {
      paddingTop: 12,
      height: '100%',
      borderRadius:12,
      backgroundColor:'rgba(0,0,0,0.4)',
      elevation: 2,
      shadowOffset: {
         width: 0,
         height: 4
      },
      shadowRadius: 6,
      shadowOpacity: 0.5
   },
   rows: {
      justifyContent: 'space-around',
      height: '25%',
      width : '100%',
      marginBottom: 12,
      paddingStart: 12,
      paddingEnd: 12,
      gap: 12,
      flex: 1,
      flexDirection: 'row',
      borderRadius: 12,
      alignItems: 'stretch',
      shadowColor: '#000000',
   },
   title: {
      textAlign:'right',
      marginEnd:12.5,
      fontSize:18,
      fontWeight:'600',
      color:'white',
      opacity: 1.0,
      textShadowColor: 'rgba(0, 0, 0, 0.75)',
      textShadowOffset: {
         width: 0,
         height: 2
      },
      textShadowRadius: 3
   },
   subtitle: {
      textAlign:'right',
      marginEnd:12.5,
      fontSize:14,
      fontWeight:'500',
      color:'white',
      opacity: 0.4
   },
   box1: {
      flex: 3,
      height: '100%',
      paddingStart:12,
      paddingEnd:12,
      paddingTop:12,
      borderTopStartRadius:12,
      borderTopEndRadius:12,
      borderRadius:12,
      alignSelf: 'flex-start',
      backgroundColor:'rgba(0,0,0,0.4)',
      elevation: 2,
      shadowOffset: {
         width: 0,
         height: 4
      },
      shadowRadius: 6,
      shadowOpacity: 0.4
   },
   subbox: {
      flex: 3,
      height: '100%',
      alignSelf: 'flex-start',
      backgroundColor:'rgba(0,0,0,0)',
   },
   avatarbox: {
      flex: 1,
      height: '100%',
      paddingStart:12,
      paddingEnd:12,
      paddingTop:12,
      borderTopStartRadius:12,
      borderTopEndRadius:12,
      borderRadius:12,
      alignSelf: 'flex-start',
      backgroundColor:'rgba(0,0,0,0)' 
   },
   name: {
      textAlign:'right',
      marginEnd:12.5,
      fontSize:36,
      fontWeight:'500',
      color:'white',
      opacity: 1.0,
      textShadowColor: 'rgba(0, 0, 0, 0.5)',
      textShadowOffset: {
         width: 0,
         height: 2
      },
      textShadowRadius: 3
   },
   username: {
      textAlign:'right',
      marginEnd:12.5,
      fontSize:18,
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
      width: 128,
      height: 128,
      borderRadius: 100,
      backgroundColor: '#FFFFFF',
      opacity: 0.5,
      justifyContent: 'flex-start',
      
   }
   
});

export default Profile;
