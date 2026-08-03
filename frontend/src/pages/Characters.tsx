import DashboardLayout from "../components/layout/DashboardLayout";
import CharactersList from "../components/characters/CharacterList";

export default function Characters(){
    return(
        <DashboardLayout>
            <h1>Guerreiros</h1>
            <CharactersList/>
        </DashboardLayout>
    );
}